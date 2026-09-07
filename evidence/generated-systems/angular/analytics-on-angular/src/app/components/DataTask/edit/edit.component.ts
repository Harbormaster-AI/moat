import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DataTaskService } from '../../../services/DataTask.service';
import { SubBaseComponent } from '../../DataTask/sub.base.component';


@Component({
    selector: 'app-edit-dataTask',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDataTaskComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DataTask';

    dataTaskForm: FormGroup;
    dataTask: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DataTaskService,
        private fb: FormBuilder
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

    
    updateDataTask(name, command, retries, Pipeline, InputDatasets, OutputDatasets, TaskType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDataTask(name, command, retries, Pipeline, InputDatasets, OutputDatasets, TaskType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDataTask']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDataTask(params['id']).subscribe(res => {
                this.dataTask = res;
            });
        });
    }
}