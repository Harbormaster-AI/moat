import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ContentCategoryService } from '../../../services/ContentCategory.service';
import { ContentCategory } from '../../../models/ContentCategory';
import { SubBaseComponent } from '../../ContentCategory/sub.base.component';

@Component({
    selector: 'app-create-contentCategory',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateContentCategoryComponent extends SubBaseComponent implements OnInit {

    title = 'Add ContentCategory';

    contentCategoryForm: FormGroup;
    contentCategory: ContentCategory;

    constructor( http: HttpClient,
        private contentCategoryService: ContentCategoryService,
        private fb: FormBuilder,
        private router: Router
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

    
    addContentCategory(code, name): void {
        this.contentCategoryService
        .addContentCategory(code, name)
            .subscribe(() => {
                this.router.navigate(['/indexContentCategory']);
            });
    }

    ngOnInit(): void {
    }
}