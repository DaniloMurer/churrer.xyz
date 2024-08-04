package xyz.churrer.house.domain.dto;

import com.fasterxml.jackson.annotation.JsonFormat;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import java.time.LocalDateTime;
import xyz.churrer.house.domain.jpa.Telemetry;

/**
 * Data transfer object class for Telemetry data
 */
@Data
@AllArgsConstructor
@NoArgsConstructor
public class TelemetryDto {

    private String countryName;
    private String countryIso;
    @JsonFormat(shape= JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd'T'HH:mm:ss")
    private LocalDateTime timestamp;

    /**
     * Convert data transfer object to jpa entity
     * @return {@link Telemetry} jpa entity
     */
    public Telemetry toTelemetry() {
        return new Telemetry(this.countryName, this.countryIso, this.timestamp);
    }

    /**
     * Create data transfer object from jpa entity
     * @param telemetry {@link Telemetry} jpa entity to convert
     * @return {@link TelemetryDto} converted data transfer object
     */
    public static TelemetryDto fromTelemetry(Telemetry telemetry) {

        return new TelemetryDto(telemetry.getCountryName(), telemetry.getCountryIso(), telemetry.getTimestamp());
    }

    @Override
    public String toString() {
        return String.format("TelemetryDto{countryName='%s', countryIso='%s', timestamp=%s}", countryName, countryIso, timestamp);
    }
}
