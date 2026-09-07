import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BuildScheduleService } from '../../../services/BuildSchedule.service';
import { BuildSchedule } from '../../../models/BuildSchedule';
import { SubBaseComponent } from '../../BuildSchedule/sub.base.component';

@Component({
    selector: 'app-create-buildSchedule',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBuildScheduleComponent extends SubBaseComponent implements OnInit {

    title = 'Add BuildSchedule';

    buildScheduleForm: FormGroup;
    buildSchedule: BuildSchedule;

    constructor( http: HttpClient,
        private buildScheduleService: BuildScheduleService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.buildScheduleForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  scheduleNumber: ['', Validators.required],
      ProductionOrders: ['', ],
      Status: ['', ]
        });
    }

    
    addBuildSchedule(scheduleNumber, ProductionOrders, Status): void {
        this.buildScheduleService
        .addBuildSchedule(scheduleNumber, ProductionOrders, Status)
            .subscribe(() => {
                this.router.navigate(['/indexBuildSchedule']);
            });
    }

    ngOnInit(): void {
    }
}