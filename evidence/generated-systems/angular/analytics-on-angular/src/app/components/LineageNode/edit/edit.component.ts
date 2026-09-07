import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LineageNodeService } from '../../../services/LineageNode.service';
import { SubBaseComponent } from '../../LineageNode/sub.base.component';


@Component({
    selector: 'app-edit-lineageNode',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLineageNodeComponent extends SubBaseComponent implements OnInit {

    title = 'Edit LineageNode';

    lineageNodeForm: FormGroup;
    lineageNode: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LineageNodeService,
        private fb: FormBuilder
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

    
    updateLineageNode(name, qualifiedName, Workspace, Inputs, Outputs, Datasets, Models, Pipelines, Dashboards, Reports, NodeType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLineageNode(name, qualifiedName, Workspace, Inputs, Outputs, Datasets, Models, Pipelines, Dashboards, Reports, NodeType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLineageNode']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLineageNode(params['id']).subscribe(res => {
                this.lineageNode = res;
            });
        });
    }
}