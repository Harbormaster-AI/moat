import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PlannedOrderService } from '../../../services/PlannedOrder.service';
import { SubBaseComponent } from '../../PlannedOrder/sub.base.component';


@Component({
    selector: 'app-edit-plannedOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPlannedOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PlannedOrder';

    plannedOrderForm: FormGroup;
    plannedOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PlannedOrderService,
        private fb: FormBuilder
) {
        super(http);
        this.plannedOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  plannedOrderNumber: ['', Validators.required],
      quantity: ['', Validators.required],
      dueDate: ['', Validators.required],
      MrpRun: ['', ],
      Item: ['', ],
      Plant: ['', ],
      OrderType: ['', ],
      Status: ['', ]
        });
    }

    
    updatePlannedOrder(plannedOrderNumber, quantity, dueDate, MrpRun, Item, Plant, OrderType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePlannedOrder(plannedOrderNumber, quantity, dueDate, MrpRun, Item, Plant, OrderType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPlannedOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPlannedOrder(params['id']).subscribe(res => {
                this.plannedOrder = res;
            });
        });
    }
}