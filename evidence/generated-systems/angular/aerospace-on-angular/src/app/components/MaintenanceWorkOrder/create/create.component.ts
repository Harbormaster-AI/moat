import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MaintenanceWorkOrderService } from '../../../services/MaintenanceWorkOrder.service';
import { MaintenanceWorkOrder } from '../../../models/MaintenanceWorkOrder';
import { SubBaseComponent } from '../../MaintenanceWorkOrder/sub.base.component';

@Component({
    selector: 'app-create-maintenanceWorkOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMaintenanceWorkOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add MaintenanceWorkOrder';

    maintenanceWorkOrderForm: FormGroup;
    maintenanceWorkOrder: MaintenanceWorkOrder;

    constructor( http: HttpClient,
        private maintenanceWorkOrderService: MaintenanceWorkOrderService,
        private fb: FormBuilder,
        private router: Router
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

    
    addMaintenanceWorkOrder(workOrderNumber, Aircraft, AirworthinessDirective, ServiceBulletin, Status): void {
        this.maintenanceWorkOrderService
        .addMaintenanceWorkOrder(workOrderNumber, Aircraft, AirworthinessDirective, ServiceBulletin, Status)
            .subscribe(() => {
                this.router.navigate(['/indexMaintenanceWorkOrder']);
            });
    }

    ngOnInit(): void {
    }
}