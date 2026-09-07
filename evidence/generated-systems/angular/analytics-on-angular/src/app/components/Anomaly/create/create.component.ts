import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AnomalyService } from '../../../services/Anomaly.service';
import { Anomaly } from '../../../models/Anomaly';
import { SubBaseComponent } from '../../Anomaly/sub.base.component';

@Component({
    selector: 'app-create-anomaly',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAnomalyComponent extends SubBaseComponent implements OnInit {

    title = 'Add Anomaly';

    anomalyForm: FormGroup;
    anomaly: Anomaly;

    constructor( http: HttpClient,
        private anomalyService: AnomalyService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAnomaly(occurredAt, details, TimeSeries, Alert, Dataset, AnomalyType, Severity): void {
        this.anomalyService
        .addAnomaly(occurredAt, details, TimeSeries, Alert, Dataset, AnomalyType, Severity)
            .subscribe(() => {
                this.router.navigate(['/indexAnomaly']);
            });
    }

    ngOnInit(): void {
    }
}