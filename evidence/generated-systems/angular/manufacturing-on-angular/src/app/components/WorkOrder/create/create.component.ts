import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { WorkOrderService } from '../../../services/WorkOrder.service';
import { WorkOrder } from '../../../models/WorkOrder';
import { SubBaseComponent } from '../../WorkOrder/sub.base.component';

@Component({
    selector: 'app-create-workOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateWorkOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add WorkOrder';

    workOrderForm: FormGroup;
    workOrder: WorkOrder;

    constructor( http: HttpClient,
        private workOrderService: WorkOrderService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.workOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  workOrderNumber: ['', Validators.required],
      plannedStart: ['', Validators.required],
      plannedEnd: ['', Validators.required],
      quantity: ['', Validators.required],
      priority: ['', Validators.required],
      Item: ['', ],
      Plant: ['', ],
      Routing: ['', ],
      Bom: ['', ],
      ProductionSchedule: ['', ],
      SalesOrder: ['', ],
      Status: ['', ]
        });
    }

    
    addWorkOrder(workOrderNumber, plannedStart, plannedEnd, quantity, priority, Item, Plant, Routing, Bom, ProductionSchedule, SalesOrder, Status): void {
        this.workOrderService
        .addWorkOrder(workOrderNumber, plannedStart, plannedEnd, quantity, priority, Item, Plant, Routing, Bom, ProductionSchedule, SalesOrder, Status)
            .subscribe(() => {
                this.router.navigate(['/indexWorkOrder']);
            });
    }

    ngOnInit(): void {
    }
}