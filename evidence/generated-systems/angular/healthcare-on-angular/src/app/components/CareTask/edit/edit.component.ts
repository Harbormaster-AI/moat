import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CareTaskService } from '../../../services/CareTask.service';
import { SubBaseComponent } from '../../CareTask/sub.base.component';


@Component({
    selector: 'app-edit-careTask',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCareTaskComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CareTask';

    careTaskForm: FormGroup;
    careTask: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CareTaskService,
        private fb: FormBuilder
) {
        super(http);
        this.careTaskForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  description: ['', Validators.required],
      dueDate: ['', Validators.required],
      CarePlan: ['', ],
      AssignedTo: ['', ],
      Encounter: ['', ],
      Status: ['', ],
      Priority: ['', ]
        });
    }

    
    updateCareTask(description, dueDate, CarePlan, AssignedTo, Encounter, Status, Priority): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCareTask(description, dueDate, CarePlan, AssignedTo, Encounter, Status, Priority, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCareTask']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCareTask(params['id']).subscribe(res => {
                this.careTask = res;
            });
        });
    }
}