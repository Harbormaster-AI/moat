import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CareTaskService } from '../../../services/CareTask.service';
import { CareTask } from '../../../models/CareTask';
import { SubBaseComponent } from '../../CareTask/sub.base.component';

@Component({
    selector: 'app-create-careTask',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCareTaskComponent extends SubBaseComponent implements OnInit {

    title = 'Add CareTask';

    careTaskForm: FormGroup;
    careTask: CareTask;

    constructor( http: HttpClient,
        private careTaskService: CareTaskService,
        private fb: FormBuilder,
        private router: Router
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

    
    addCareTask(description, dueDate, CarePlan, AssignedTo, Encounter, Status, Priority): void {
        this.careTaskService
        .addCareTask(description, dueDate, CarePlan, AssignedTo, Encounter, Status, Priority)
            .subscribe(() => {
                this.router.navigate(['/indexCareTask']);
            });
    }

    ngOnInit(): void {
    }
}