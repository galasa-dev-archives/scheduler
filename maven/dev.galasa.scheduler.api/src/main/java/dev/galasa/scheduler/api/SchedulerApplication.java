/*
 * Copyright contributors to the Galasa project
 */
package dev.galasa.scheduler.api;

import java.util.HashSet;
import java.util.Set;

import javax.ws.rs.ApplicationPath;
import javax.ws.rs.core.Application;

import dev.galasa.scheduler.api.status.StatusRestImpl;

@ApplicationPath("")
public class SchedulerApplication extends Application {
	
	@Override
	public Set<Class<?>> getClasses() {
		
		HashSet<Class<?>> classes = new HashSet<>();
		classes.add(StatusRestImpl.class);
		
		return classes;
	}

}
