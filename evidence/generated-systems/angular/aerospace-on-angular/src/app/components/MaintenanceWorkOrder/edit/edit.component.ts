import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MaintenanceWorkOrderService } from '../../../services/MaintenanceWorkOrder.service';
import { SubBaseComponent } from '../../MaintenanceWorkOrder/sub.base.component';


@Component({
    selector: 'app-edit-maintenanceWorkOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMaintenanceWorkOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit MaintenanceWorkOrder';

    maintenanceWorkOrderForm: FormGroup;
    maintenanceWorkOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MaintenanceWorkOrderService,
        private fb: FormBuilder
) {
        super(http);
        this.maintenanceWorkOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  workOrderNumber: ['', Validators.required],
      Aircraft: ['', ],
      AirworthinessDirective: ['', ],
      ServiceBulletin: ['', ],
      Status: ['', ]
        });
    }

    
    updateMaintenanceWorkOrder(workOrderNumber, Aircraft, AirworthinessDirective, ServiceBulletin, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMaintenanceWorkOrder(workOrderNumber, Aircraft, AirworthinessDirective, ServiceBulletin, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMaintenanceWorkOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMaintenanceWorkOrder(params['id']).subscribe(res => {
                this.maintenanceWorkOrder = res;
            });
        });
    }
}