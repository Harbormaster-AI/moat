import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TimeSeriesService } from '../../../services/TimeSeries.service';
import { TimeSeries } from '../../../models/TimeSeries';
import { SubBaseComponent } from '../../TimeSeries/sub.base.component';

@Component({
    selector: 'app-create-timeSeries',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTimeSeriesComponent extends SubBaseComponent implements OnInit {

    title = 'Add TimeSeries';

    timeSeriesForm: FormGroup;
    timeSeries: TimeSeries;

    constructor( http: HttpClient,
        private timeSeriesService: TimeSeriesService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.timeSeriesForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      timezone: ['', Validators.required],
      Datasets: ['', ],
      Forecasts: ['', ],
      Anomalies: ['', ],
      Granularity: ['', ]
        });
    }

    
    addTimeSeries(name, timezone, Datasets, Forecasts, Anomalies, Granularity): void {
        this.timeSeriesService
        .addTimeSeries(name, timezone, Datasets, Forecasts, Anomalies, Granularity)
            .subscribe(() => {
                this.router.navigate(['/indexTimeSeries']);
            });
    }

    ngOnInit(): void {
    }
}