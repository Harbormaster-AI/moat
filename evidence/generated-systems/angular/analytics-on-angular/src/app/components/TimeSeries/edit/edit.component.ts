import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TimeSeriesService } from '../../../services/TimeSeries.service';
import { SubBaseComponent } from '../../TimeSeries/sub.base.component';


@Component({
    selector: 'app-edit-timeSeries',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTimeSeriesComponent extends SubBaseComponent implements OnInit {

    title = 'Edit TimeSeries';

    timeSeriesForm: FormGroup;
    timeSeries: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TimeSeriesService,
        private fb: FormBuilder
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

    
    updateTimeSeries(name, timezone, Datasets, Forecasts, Anomalies, Granularity): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTimeSeries(name, timezone, Datasets, Forecasts, Anomalies, Granularity, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTimeSeries']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTimeSeries(params['id']).subscribe(res => {
                this.timeSeries = res;
            });
        });
    }
}