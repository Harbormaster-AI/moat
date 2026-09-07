import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BuildScheduleService } from '../../../services/BuildSchedule.service';
import { SubBaseComponent } from '../../BuildSchedule/sub.base.component';


@Component({
    selector: 'app-edit-buildSchedule',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBuildScheduleComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BuildSchedule';

    buildScheduleForm: FormGroup;
    buildSchedule: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BuildScheduleService,
        private fb: FormBuilder
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

    
    updateBuildSchedule(scheduleNumber, ProductionOrders, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBuildSchedule(scheduleNumber, ProductionOrders, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBuildSchedule']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBuildSchedule(params['id']).subscribe(res => {
                this.buildSchedule = res;
            });
        });
    }
}