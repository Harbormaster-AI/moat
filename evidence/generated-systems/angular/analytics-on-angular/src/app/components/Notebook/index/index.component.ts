
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NotebookService } from '../../../services/Notebook.service';
import { Notebook } from '../../../models/Notebook';

@Component({
    selector: 'app-index-notebook',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexNotebookComponent implements OnInit {

    notebooks: Notebook[] = [];

    constructor(
        private router: Router,
        private service: NotebookService
) {}

    ngOnInit(): void {
        this.getNotebooks();
}

    getNotebooks(): void {
        this.service.getNotebooks().subscribe((res) => {
        this.notebooks = res;
    });
}

    deleteNotebook(id: any): void {
        this.service.deleteNotebook(id)
            .subscribe(() => {
                this.getNotebooks();
            });
    }
}