import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PerformanceMetricService } from '../../../services/PerformanceMetric.service';
import { SubBaseComponent } from '../../PerformanceMetric/sub.base.component';


@Component({
    selector: 'app-edit-performanceMetric',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPerformanceMetricComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PerformanceMetric';

    performanceMetricForm: FormGroup;
    performanceMetric: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PerformanceMetricService,
        private fb: FormBuilder
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

    
    updatePerformanceMetric(date, value, AdAccount, Campaign, LineItem, Placement, CreativeAsset, MetricType): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePerformanceMetric(date, value, AdAccount, Campaign, LineItem, Placement, CreativeAsset, MetricType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPerformanceMetric']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPerformanceMetric(params['id']).subscribe(res => {
                this.performanceMetric = res;
            });
        });
    }
}