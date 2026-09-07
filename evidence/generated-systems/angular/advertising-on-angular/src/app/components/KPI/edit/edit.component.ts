import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { KPIService } from '../../../services/KPI.service';
import { SubBaseComponent } from '../../KPI/sub.base.component';


@Component({
    selector: 'app-edit-kPI',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditKPIComponent extends SubBaseComponent implements OnInit {

    title = 'Edit KPI';

    kPIForm: FormGroup;
    kPI: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: KPIService,
        private fb: FormBuilder
) {
        super(http);
        this.kPIForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  targetValue: ['', Validators.required],
      Campaign: ['', ],
      MetricType: ['', ]
        });
    }

    
    updateKPI(targetValue, Campaign, MetricType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateKPI(targetValue, Campaign, MetricType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexKPI']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getKPI(params['id']).subscribe(res => {
                this.kPI = res;
            });
        });
    }
}