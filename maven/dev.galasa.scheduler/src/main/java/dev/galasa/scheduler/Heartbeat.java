package dev.galasa.scheduler;

import java.time.Instant;

import javax.persistence.EntityManager;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import dev.galasa.scheduler.jpa.StatusEntity;

public class Heartbeat implements Runnable {

	private final Scheduler scheduler;
	
	private final Logger logger = LogManager.getLogger(Heartbeat.class);


	public Heartbeat(Scheduler scheduler) {
		this.scheduler = scheduler;
	}

	@Override
	public void run() {
		System.out.println("ee1k");
		try {
			EntityManager em = this.scheduler.getEntityManager();
			em.getTransaction().begin();
			System.out.println("ee2k");

			StatusEntity statusRecord = em.find(StatusEntity.class, "scheduler");
			if (statusRecord == null) {
				statusRecord = new StatusEntity();
				statusRecord.setId("scheduler");
			}
			System.out.println("ee3k");
			
			statusRecord.setSummary("Idle");
			statusRecord.setTimestamp(Instant.now());
			em.persist(statusRecord);
			
			em.getTransaction().commit();

			logger.trace("Heartbeat complete");
		} catch(Exception e) {
			e.printStackTrace();
		}

	}

}
