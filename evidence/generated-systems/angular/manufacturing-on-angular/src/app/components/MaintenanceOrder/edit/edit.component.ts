import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MaintenanceOrderService } from '../../../services/MaintenanceOrder.service';
import { SubBaseComponent } from '../../MaintenanceOrder/sub.base.component';


@Component({
    selector: 'app-edit-maintenanceOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMaintenanceOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit MaintenanceOrder';

    maintenanceOrderForm: FormGroup;
    maintenanceOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MaintenanceOrderService,
        private fb: FormBuilder
) {
        super(http);
        this.maintenanceOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  orderNumber: ['', Validators.required],
      priority: ['', Validators.required],
      requestedDate: ['', Validators.required],
      completionDate: ['', Validators.required],
      Asset: ['', ],
      Plan: ['', ],
      WorkCenter: ['', ],
      Status: ['', ]
        });
    }

    
    updateMaintenanceOrder(orderNumber, priority, requestedDate, completionDate, Asset, Plan, WorkCenter, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMaintenanceOrder(orderNumber, priority, requestedDate, completionDate, Asset, Plan, WorkCenter, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMaintenanceOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMaintenanceOrder(params['id']).subscribe(res => {
                this.maintenanceOrder = res;
            });
        });
    }
}