package dev.galasa.scheduler.api;

import javax.enterprise.context.ApplicationScoped;
import javax.enterprise.context.RequestScoped;
import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;

import org.eclipse.microprofile.graphql.GraphQLApi;
import org.eclipse.microprofile.graphql.Query;

import dev.galasa.scheduler.api.model.Status;
import dev.galasa.scheduler.jpa.StatusEntity;

@GraphQLApi
@ApplicationScoped
public class StatusGraph {
	
	@PersistenceContext(name = "jpa-unit")
	private EntityManager em;

	@Query
	public Status getServerStatus() {
		Status status = new Status();

		StatusEntity jpaStatus = null;
		try {
			jpaStatus = em.find(StatusEntity.class, "scheduler");
			if (jpaStatus == null) {
				status.apiReport = "Ok";
				status.schedulerReport ="Not started";
			} else {
				status.apiReport = "Ok";
				status.schedulerReport = jpaStatus.getSummary();
			}
		} catch(Exception e) {
			status.apiReport = "Database is not initialised";
			status.schedulerReport ="Unknown";
			e.printStackTrace();
		}

		return status;
		
	}

}
