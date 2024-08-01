package xyz.churrer.house.resource;

import jakarta.inject.Inject;
import jakarta.ws.rs.POST;
import jakarta.ws.rs.Path;
import xyz.churrer.house.domain.jpa.Telemetry;
import xyz.churrer.house.service.TelemetryService;

import java.time.LocalDateTime;

@Path("/api/telemetry")
public class TelemetryResource {

    @Inject
    TelemetryService telemetryService;

    @POST
    public String mockData() {
        var telemetry = new Telemetry("", "", LocalDateTime.now());
        this.telemetryService.persist(telemetry);
        return "";
    }
}
