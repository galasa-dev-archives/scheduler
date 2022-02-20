/*
 * Copyright contributors to the Galasa project
 */
package dev.galasa.scheduler.api.status;

import javax.enterprise.context.RequestScoped;
import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;

import dev.galasa.scheduler.api.openapi.StatusApi;
import dev.galasa.scheduler.api.openapi.model.Status;
import dev.galasa.scheduler.jpa.StatusEntity;


@RequestScoped
public class StatusRestImpl implements StatusApi {

	@PersistenceContext(name = "jpa-unit")
	private EntityManager em;

	@Override
	public dev.galasa.scheduler.api.openapi.model.Status getStatus() {
		
		Status status = new Status();

		StatusEntity jpaStatus = null;
		try {
			jpaStatus = em.find(StatusEntity.class, "scheduler");
			status.apiReport("Ok");
			status.schedulerReport(jpaStatus.getSummary());
		} catch(Exception e) {
			status.apiReport("Database is not initialised");
			status.schedulerReport("Unknown");
		}

		return status;
	}

}
