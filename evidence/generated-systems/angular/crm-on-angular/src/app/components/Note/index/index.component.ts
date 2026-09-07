
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NoteService } from '../../../services/Note.service';
import { Note } from '../../../models/Note';

@Component({
    selector: 'app-index-note',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexNoteComponent implements OnInit {

    notes: Note[] = [];

    constructor(
        private router: Router,
        private service: NoteService
) {}

    ngOnInit(): void {
        this.getNotes();
}

    getNotes(): void {
        this.service.getNotes().subscribe((res) => {
        this.notes = res;
    });
}

    deleteNote(id: any): void {
        this.service.deleteNote(id)
            .subscribe(() => {
                this.getNotes();
            });
    }
}