import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ContentCategoryService } from '../../../services/ContentCategory.service';
import { SubBaseComponent } from '../../ContentCategory/sub.base.component';


@Component({
    selector: 'app-edit-contentCategory',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditContentCategoryComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ContentCategory';

    contentCategoryForm: FormGroup;
    contentCategory: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ContentCategoryService,
        private fb: FormBuilder
) {
        super(http);
        this.contentCategoryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  code: ['', Validators.required],
      name: ['', Validators.required]
        });
    }

    
    updateContentCategory(code, name): void {
        this.route.params.subscribe((params) => {

                        this.service.updateContentCategory(code, name, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexContentCategory']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getContentCategory(params['id']).subscribe(res => {
                this.contentCategory = res;
            });
        });
    }
}