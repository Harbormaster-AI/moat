import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MaintenancePlanService } from '../../../services/MaintenancePlan.service';
import { SubBaseComponent } from '../../MaintenancePlan/sub.base.component';


@Component({
    selector: 'app-edit-maintenancePlan',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMaintenancePlanComponent extends SubBaseComponent implements OnInit {

    title = 'Edit MaintenancePlan';

    maintenancePlanForm: FormGroup;
    maintenancePlan: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MaintenancePlanService,
        private fb: FormBuilder
) {
        super(http);
        this.maintenancePlanForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  planNumber: ['', Validators.required],
      interval: ['', Validators.required],
      lastServiceDate: ['', Validators.required],
      Asset: ['', ],
      MaintenanceOrders: ['', ],
      Strategy: ['', ]
        });
    }

    
    updateMaintenancePlan(planNumber, interval, lastServiceDate, Asset, MaintenanceOrders, Strategy): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMaintenancePlan(planNumber, interval, lastServiceDate, Asset, MaintenanceOrders, Strategy, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMaintenancePlan']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMaintenancePlan(params['id']).subscribe(res => {
                this.maintenancePlan = res;
            });
        });
    }
}