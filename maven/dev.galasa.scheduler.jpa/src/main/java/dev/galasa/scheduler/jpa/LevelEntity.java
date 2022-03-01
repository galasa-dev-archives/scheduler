package dev.galasa.scheduler.jpa;

import java.io.Serializable;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.Lob;
import javax.persistence.ManyToOne;
import javax.persistence.Table;
import javax.persistence.Version;

@Entity
@Table(name = "level")
public class LevelEntity implements Serializable {
	private static final long serialVersionUID = 1L;
	
	@Id
	@Column(name = "id", length = 10)
	private String id;
	
	@ManyToOne(fetch = FetchType.LAZY)
	@JoinColumn(name = "app_id")
	private AppEntity app;
	
	@Lob
	@Column(name = "configuration")
	private String configuration;
	
	@Version
	private long version;
	
	public LevelEntity() {
	}

	public String getId() {
		return id;
	}

	public void setId(String id) {
		this.id = id;
	}

	public String getConfiguration() {
		return configuration;
	}

	public void setConfiguration(String configuration) {
		this.configuration = configuration;
	}
	
	public AppEntity getApp() {
		return this.app;
	}

}
