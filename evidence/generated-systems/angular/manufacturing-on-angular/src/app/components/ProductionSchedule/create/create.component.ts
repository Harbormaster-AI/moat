import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ProductionScheduleService } from '../../../services/ProductionSchedule.service';
import { ProductionSchedule } from '../../../models/ProductionSchedule';
import { SubBaseComponent } from '../../ProductionSchedule/sub.base.component';

@Component({
    selector: 'app-create-productionSchedule',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateProductionScheduleComponent extends SubBaseComponent implements OnInit {

    title = 'Add ProductionSchedule';

    productionScheduleForm: FormGroup;
    productionSchedule: ProductionSchedule;

    constructor( http: HttpClient,
        private productionScheduleService: ProductionScheduleService,
        private fb: FormBuilder,
        private router: Router
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

    
    addProductionSchedule(scheduleNumber, horizonStart, horizonEnd, Plant, WorkOrders, Status): void {
        this.productionScheduleService
        .addProductionSchedule(scheduleNumber, horizonStart, horizonEnd, Plant, WorkOrders, Status)
            .subscribe(() => {
                this.router.navigate(['/indexProductionSchedule']);
            });
    }

    ngOnInit(): void {
    }
}