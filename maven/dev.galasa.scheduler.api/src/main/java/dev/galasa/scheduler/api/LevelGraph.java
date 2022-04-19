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
import org.eclipse.microprofile.graphql.Source;

import dev.galasa.scheduler.api.model.App;
import dev.galasa.scheduler.api.model.Level;

@GraphQLApi
@ApplicationScoped
public class LevelGraph {
	
	@PersistenceContext(name = "jpa-unit")
	private EntityManager em;

	@Query
	@Description("Retrieve all Levels")
	public List<Level> getAllLevels() {
		return new ArrayList<>();
	}
	
	@Query 
	@Description("Retrieve Levels for App")
	public List<Level> getLevels(@Source App app) {
		
		ArrayList<Level> levels = new ArrayList<>();
		
		Level demo = new Level();
		demo.id = app.id + "-demo";
		demo.name = "Demo";
		demo.app = app.id;
		demo.configuration = "";
		
		levels.add(demo);
		
		return levels;
	}

	@Query 
	@Description("Retrieve an Level by ID")
	public Level getLevelById(String id) {
		
		Level demo = new Level();
		demo.id = id;
		demo.name = "Demo";
		demo.app = "its id";
		demo.configuration = "";
		
		return demo;
	}

	@Mutation
	@Description("Create an Level")
	public Level createLevel(Level Level) {
		return null;
	}
	
	@Mutation
	@Description("Delete an Level")
	public int deleteLevel(String id) {
		return 0;
	}
	
	@Mutation
	@Description("Update an Level")
	public Level updateLevel(Level Level) {
		return null;
	}
}
