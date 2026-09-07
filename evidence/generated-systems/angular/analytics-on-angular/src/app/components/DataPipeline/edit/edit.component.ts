import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DataPipelineService } from '../../../services/DataPipeline.service';
import { SubBaseComponent } from '../../DataPipeline/sub.base.component';


@Component({
    selector: 'app-edit-dataPipeline',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDataPipelineComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DataPipeline';

    dataPipelineForm: FormGroup;
    dataPipeline: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DataPipelineService,
        private fb: FormBuilder
) {
        super(http);
        this.dataPipelineForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      schedule: ['', Validators.required],
      Workspace: ['', ],
      Tasks: ['', ],
      Sources: ['', ],
      Outputs: ['', ],
      LineageNode: ['', ],
      TriggerType: ['', ],
      Status: ['', ]
        });
    }

    
    updateDataPipeline(name, schedule, Workspace, Tasks, Sources, Outputs, LineageNode, TriggerType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDataPipeline(name, schedule, Workspace, Tasks, Sources, Outputs, LineageNode, TriggerType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDataPipeline']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDataPipeline(params['id']).subscribe(res => {
                this.dataPipeline = res;
            });
        });
    }
}