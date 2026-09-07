import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { NotebookService } from '../../../services/Notebook.service';
import { Notebook } from '../../../models/Notebook';
import { SubBaseComponent } from '../../Notebook/sub.base.component';

@Component({
    selector: 'app-create-notebook',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateNotebookComponent extends SubBaseComponent implements OnInit {

    title = 'Add Notebook';

    notebookForm: FormGroup;
    notebook: Notebook;

    constructor( http: HttpClient,
        private notebookService: NotebookService,
        private fb: FormBuilder,
        private router: Router
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

    
    addNotebook(title, repository, Workspace, Datasets, Experiments, Queries, Language): void {
        this.notebookService
        .addNotebook(title, repository, Workspace, Datasets, Experiments, Queries, Language)
            .subscribe(() => {
                this.router.navigate(['/indexNotebook']);
            });
    }

    ngOnInit(): void {
    }
}