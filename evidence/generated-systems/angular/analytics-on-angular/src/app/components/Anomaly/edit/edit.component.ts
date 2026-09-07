import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AnomalyService } from '../../../services/Anomaly.service';
import { SubBaseComponent } from '../../Anomaly/sub.base.component';


@Component({
    selector: 'app-edit-anomaly',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAnomalyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Anomaly';

    anomalyForm: FormGroup;
    anomaly: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AnomalyService,
        private fb: FormBuilder
) {
        super(http);
        this.anomalyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  occurredAt: ['', Validators.required],
      details: ['', Validators.required],
      TimeSeries: ['', ],
      Alert: ['', ],
      Dataset: ['', ],
      AnomalyType: ['', ],
      Severity: ['', ]
        });
    }

    
    updateAnomaly(occurredAt, details, TimeSeries, Alert, Dataset, AnomalyType, Severity): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAnomaly(occurredAt, details, TimeSeries, Alert, Dataset, AnomalyType, Severity, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAnomaly']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAnomaly(params['id']).subscribe(res => {
                this.anomaly = res;
            });
        });
    }
}