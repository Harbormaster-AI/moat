import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { NotebookService } from '../../../services/Notebook.service';
import { SubBaseComponent } from '../../Notebook/sub.base.component';


@Component({
    selector: 'app-edit-notebook',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditNotebookComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Notebook';

    notebookForm: FormGroup;
    notebook: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: NotebookService,
        private fb: FormBuilder
) {
        super(http);
        this.notebookForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      repository: ['', Validators.required],
      Workspace: ['', ],
      Datasets: ['', ],
      Experiments: ['', ],
      Queries: ['', ],
      Language: ['', ]
        });
    }

    
    updateNotebook(title, repository, Workspace, Datasets, Experiments, Queries, Language): void {
        this.route.params.subscribe((params) => {

                        this.service.updateNotebook(title, repository, Workspace, Datasets, Experiments, Queries, Language, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexNotebook']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getNotebook(params['id']).subscribe(res => {
                this.notebook = res;
            });
        });
    }
}