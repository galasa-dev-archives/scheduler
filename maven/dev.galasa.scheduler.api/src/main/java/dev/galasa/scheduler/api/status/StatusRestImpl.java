/*
 * Copyright contributors to the Galasa project
 */
package dev.galasa.scheduler.api.status;

import dev.galasa.scheduler.api.openapi.StatusApi;
import dev.galasa.scheduler.api.openapi.model.Status;

public class StatusRestImpl implements StatusApi {

	@Override
	public Status getStatus() {
		Status status = new Status();
		status.setApiReport("Hello");
		status.setSchedulerReport("There");
		return status;
	}

}
