import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { WorkOrderService } from '../../../services/WorkOrder.service';
import { SubBaseComponent } from '../../WorkOrder/sub.base.component';


@Component({
    selector: 'app-edit-workOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditWorkOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit WorkOrder';

    workOrderForm: FormGroup;
    workOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: WorkOrderService,
        private fb: FormBuilder
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

    
    updateWorkOrder(workOrderNumber, plannedStart, plannedEnd, quantity, priority, Item, Plant, Routing, Bom, ProductionSchedule, SalesOrder, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateWorkOrder(workOrderNumber, plannedStart, plannedEnd, quantity, priority, Item, Plant, Routing, Bom, ProductionSchedule, SalesOrder, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexWorkOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getWorkOrder(params['id']).subscribe(res => {
                this.workOrder = res;
            });
        });
    }
}