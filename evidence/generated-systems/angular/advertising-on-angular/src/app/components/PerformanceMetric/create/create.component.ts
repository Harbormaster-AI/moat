import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PerformanceMetricService } from '../../../services/PerformanceMetric.service';
import { PerformanceMetric } from '../../../models/PerformanceMetric';
import { SubBaseComponent } from '../../PerformanceMetric/sub.base.component';

@Component({
    selector: 'app-create-performanceMetric',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePerformanceMetricComponent extends SubBaseComponent implements OnInit {

    title = 'Add PerformanceMetric';

    performanceMetricForm: FormGroup;
    performanceMetric: PerformanceMetric;

    constructor( http: HttpClient,
        private performanceMetricService: PerformanceMetricService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.performanceMetricForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  date: ['', Validators.required],
      value: ['', Validators.required],
      AdAccount: ['', ],
      Campaign: ['', ],
      LineItem: ['', ],
      Placement: ['', ],
      CreativeAsset: ['', ],
      MetricType: ['', ]
        });
    }

    
    addPerformanceMetric(date, value, AdAccount, Campaign, LineItem, Placement, CreativeAsset, MetricType): void {
        this.performanceMetricService
        .addPerformanceMetric(date, value, AdAccount, Campaign, LineItem, Placement, CreativeAsset, MetricType)
            .subscribe(() => {
                this.router.navigate(['/indexPerformanceMetric']);
            });
    }

    ngOnInit(): void {
    }
}