import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DataPipelineService } from '../../../services/DataPipeline.service';
import { DataPipeline } from '../../../models/DataPipeline';
import { SubBaseComponent } from '../../DataPipeline/sub.base.component';

@Component({
    selector: 'app-create-dataPipeline',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDataPipelineComponent extends SubBaseComponent implements OnInit {

    title = 'Add DataPipeline';

    dataPipelineForm: FormGroup;
    dataPipeline: DataPipeline;

    constructor( http: HttpClient,
        private dataPipelineService: DataPipelineService,
        private fb: FormBuilder,
        private router: Router
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

    
    addDataPipeline(name, schedule, Workspace, Tasks, Sources, Outputs, LineageNode, TriggerType, Status): void {
        this.dataPipelineService
        .addDataPipeline(name, schedule, Workspace, Tasks, Sources, Outputs, LineageNode, TriggerType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexDataPipeline']);
            });
    }

    ngOnInit(): void {
    }
}