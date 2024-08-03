package xyz.churrer.house.domain.dto;

import com.fasterxml.jackson.annotation.JsonFormat;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import xyz.churrer.house.domain.jpa.Telemetry;

import java.time.LocalDateTime;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class TelemetryDto {
    private String countryName;
    private String countryIso;
    @JsonFormat(shape= JsonFormat.Shape.STRING, pattern = "yyyy-MM-dd'T'HH:mm:ss")
    private LocalDateTime timestamp;

    public Telemetry toTelemetry() {
        return new Telemetry(this.countryName, this.countryIso, this.timestamp);
    }

    public static TelemetryDto fromTelemetry(Telemetry telemetry) {

        return new TelemetryDto(telemetry.getCountryName(), telemetry.getCountryIso(), telemetry.getTimestamp());
    }

    @Override
    public String toString() {
        return String.format("TelemetryDto{countryName='%s', countryIso='%s', timestamp=%s}", countryName, countryIso, timestamp);
    }
}
