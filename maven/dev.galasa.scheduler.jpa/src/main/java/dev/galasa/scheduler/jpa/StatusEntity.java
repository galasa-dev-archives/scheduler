package dev.galasa.scheduler.jpa;

import java.io.Serializable;
import java.time.Instant;
import java.util.Date;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import javax.persistence.Temporal;
import javax.persistence.TemporalType;

@Entity
@Table(name = "status")
public class StatusEntity implements Serializable {
	private static final long serialVersionUID = 1L;
	
	@Id
	@Column(name = "id", length = 10)
	private String id;
	
	@Column(name = "summary", length=256)
	private String summary;
	
	// Must be UTC
	@Column(name = "timestamp")
	@Temporal(TemporalType.TIMESTAMP)
	private Date timestamp;
	
	public StatusEntity() {
	}

	public String getId() {
		return id;
	}

	public void setId(String id) {
		this.id = id;
	}

	public String getSummary() {
		return summary;
	}

	public void setSummary(String summary) {
		this.summary = summary;
	}

	public Instant getTimestamp() {
		return timestamp.toInstant();
	}

	public void setTimestamp(Instant timestamp) {
		this.timestamp = Date.from(timestamp);
	}

}
