package dev.galasa.scheduler;

import java.util.Properties;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.ScheduledThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.Persistence;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class Scheduler {
	
	private EntityManager em;
	private ScheduledExecutorService scheduledExecutorService;
	
	private Heartbeat heartbeat = new Heartbeat(this);
	
	private final Logger logger = LogManager.getLogger(Scheduler.class);
	
	private boolean shutdown = false;

	private void run() {
		logger.info("Startup");
		
        Properties emProperties = new Properties();
        String envUrl = System.getenv("POSTGRESQL_URL");
        if (envUrl != null) {
        	emProperties.put("javax.persistence.jdbc.url", envUrl);
        	logger.info("Connecting to database '" + envUrl + "'");
        }
        EntityManagerFactory emf = Persistence.createEntityManagerFactory("jpa-unit",emProperties);
        this.em = emf.createEntityManager(emProperties);
        
        System.out.println("boo");
        
        this.scheduledExecutorService = new ScheduledThreadPoolExecutor(2);
        this.scheduledExecutorService.scheduleWithFixedDelay(this.heartbeat, 0, 5, TimeUnit.SECONDS);
        
        while(!shutdown) {
        	try {
				Thread.sleep(100);
			} catch (InterruptedException e) {
				Thread.currentThread().interrupt();
				return;
			}
        }

        logger.info("Shutting down");
        
        this.scheduledExecutorService.shutdown();
        try {
			this.scheduledExecutorService.awaitTermination(10, TimeUnit.SECONDS);
		} catch (InterruptedException e) {
			Thread.currentThread().interrupt();
		}
        
        em.close();

        logger.info("Shutdown complete");
	}

	public static void main(String[] args) {
		new Scheduler().run();
	}

	public EntityManager getEntityManager() {
		return this.em;
	}
	
	public class ShutdownHook extends Thread {
		
		@Override
		public void run() {
			Scheduler.this.logger.info("Shutdown requested");
			Scheduler.this.shutdown = true;
		}
		
	}

}
