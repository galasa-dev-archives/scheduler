package dev.galasa.scheduler.api;

import java.util.ArrayList;
import java.util.List;

import javax.enterprise.context.ApplicationScoped;
import javax.enterprise.context.RequestScoped;
import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;

import org.eclipse.microprofile.graphql.Description;
import org.eclipse.microprofile.graphql.GraphQLApi;
import org.eclipse.microprofile.graphql.Mutation;
import org.eclipse.microprofile.graphql.Query;

import dev.galasa.scheduler.api.model.App;

@GraphQLApi
@ApplicationScoped
public class AppGraph {
	
	@PersistenceContext(name = "jpa-unit")
	private EntityManager em;

	@Query
	@Description("Retrieve all Apps")
	public List<App> getAllApps() {
		return new ArrayList<>();
	}
	
	@Query 
	@Description("Retrieve an App by ID")
	public App getApp(String id) {
		
		App demo = new App();
		demo.id = id;
		demo.name = "Demo";
		demo.configuration = "";
		
		return demo;
	}

	@Mutation
	@Description("Create an App")
	public App createApp(App app) {
		return null;
	}
	
	@Mutation
	@Description("Delete an App")
	public int deleteApp(String id) {
		return 0;
	}
	
	@Mutation
	@Description("Update an App")
	public App updateApp(App app) {
		return null;
	}
}
