import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { NoteService } from '../../../services/Note.service';
import { SubBaseComponent } from '../../Note/sub.base.component';


@Component({
    selector: 'app-edit-note',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditNoteComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Note';

    noteForm: FormGroup;
    note: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: NoteService,
        private fb: FormBuilder
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

    
    updateNote(title, content, createdAt, updatedAt, Organization, Owner, Account, Contact, Opportunity, Case, Lead): void {
        this.route.params.subscribe((params) => {

                        this.service.updateNote(title, content, createdAt, updatedAt, Organization, Owner, Account, Contact, Opportunity, Case, Lead, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexNote']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getNote(params['id']).subscribe(res => {
                this.note = res;
            });
        });
    }
}