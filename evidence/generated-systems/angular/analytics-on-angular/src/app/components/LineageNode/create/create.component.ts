import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LineageNodeService } from '../../../services/LineageNode.service';
import { LineageNode } from '../../../models/LineageNode';
import { SubBaseComponent } from '../../LineageNode/sub.base.component';

@Component({
    selector: 'app-create-lineageNode',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLineageNodeComponent extends SubBaseComponent implements OnInit {

    title = 'Add LineageNode';

    lineageNodeForm: FormGroup;
    lineageNode: LineageNode;

    constructor( http: HttpClient,
        private lineageNodeService: LineageNodeService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.lineageNodeForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      qualifiedName: ['', Validators.required],
      Workspace: ['', ],
      Inputs: ['', ],
      Outputs: ['', ],
      Datasets: ['', ],
      Models: ['', ],
      Pipelines: ['', ],
      Dashboards: ['', ],
      Reports: ['', ],
      NodeType: ['', ]
        });
    }

    
    addLineageNode(name, qualifiedName, Workspace, Inputs, Outputs, Datasets, Models, Pipelines, Dashboards, Reports, NodeType): void {
        this.lineageNodeService
        .addLineageNode(name, qualifiedName, Workspace, Inputs, Outputs, Datasets, Models, Pipelines, Dashboards, Reports, NodeType)
            .subscribe(() => {
                this.router.navigate(['/indexLineageNode']);
            });
    }

    ngOnInit(): void {
    }
}