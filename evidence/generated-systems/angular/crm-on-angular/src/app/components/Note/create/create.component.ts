import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { NoteService } from '../../../services/Note.service';
import { Note } from '../../../models/Note';
import { SubBaseComponent } from '../../Note/sub.base.component';

@Component({
    selector: 'app-create-note',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateNoteComponent extends SubBaseComponent implements OnInit {

    title = 'Add Note';

    noteForm: FormGroup;
    note: Note;

    constructor( http: HttpClient,
        private noteService: NoteService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.noteForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      content: ['', Validators.required],
      createdAt: ['', Validators.required],
      updatedAt: ['', Validators.required],
      Organization: ['', ],
      Owner: ['', ],
      Account: ['', ],
      Contact: ['', ],
      Opportunity: ['', ],
      Case: ['', ],
      Lead: ['', ]
        });
    }

    
    addNote(title, content, createdAt, updatedAt, Organization, Owner, Account, Contact, Opportunity, Case, Lead): void {
        this.noteService
        .addNote(title, content, createdAt, updatedAt, Organization, Owner, Account, Contact, Opportunity, Case, Lead)
            .subscribe(() => {
                this.router.navigate(['/indexNote']);
            });
    }

    ngOnInit(): void {
    }
}