import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DataTaskService } from '../../../services/DataTask.service';
import { DataTask } from '../../../models/DataTask';
import { SubBaseComponent } from '../../DataTask/sub.base.component';

@Component({
    selector: 'app-create-dataTask',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDataTaskComponent extends SubBaseComponent implements OnInit {

    title = 'Add DataTask';

    dataTaskForm: FormGroup;
    dataTask: DataTask;

    constructor( http: HttpClient,
        private dataTaskService: DataTaskService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.dataTaskForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      command: ['', Validators.required],
      retries: ['', Validators.required],
      Pipeline: ['', ],
      InputDatasets: ['', ],
      OutputDatasets: ['', ],
      TaskType: ['', ]
        });
    }

    
    addDataTask(name, command, retries, Pipeline, InputDatasets, OutputDatasets, TaskType): void {
        this.dataTaskService
        .addDataTask(name, command, retries, Pipeline, InputDatasets, OutputDatasets, TaskType)
            .subscribe(() => {
                this.router.navigate(['/indexDataTask']);
            });
    }

    ngOnInit(): void {
    }
}