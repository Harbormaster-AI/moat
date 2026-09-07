import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ProductionScheduleService } from '../../../services/ProductionSchedule.service';
import { SubBaseComponent } from '../../ProductionSchedule/sub.base.component';


@Component({
    selector: 'app-edit-productionSchedule',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditProductionScheduleComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ProductionSchedule';

    productionScheduleForm: FormGroup;
    productionSchedule: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ProductionScheduleService,
        private fb: FormBuilder
) {
        super(http);
        this.productionScheduleForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  scheduleNumber: ['', Validators.required],
      horizonStart: ['', Validators.required],
      horizonEnd: ['', Validators.required],
      Plant: ['', ],
      WorkOrders: ['', ],
      Status: ['', ]
        });
    }

    
    updateProductionSchedule(scheduleNumber, horizonStart, horizonEnd, Plant, WorkOrders, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateProductionSchedule(scheduleNumber, horizonStart, horizonEnd, Plant, WorkOrders, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexProductionSchedule']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getProductionSchedule(params['id']).subscribe(res => {
                this.productionSchedule = res;
            });
        });
    }
}