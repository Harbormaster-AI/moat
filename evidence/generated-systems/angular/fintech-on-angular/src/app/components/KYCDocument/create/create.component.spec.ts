
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateKYCDocumentComponent } from './create.component';
import { KYCDocumentService } from '../../../services/KYCDocument.service';
import { Router } from '@angular/router';

describe('CreateKYCDocumentComponent', () => {
  let component: CreateKYCDocumentComponent;
  let fixture: ComponentFixture<CreateKYCDocumentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateKYCDocumentComponent
      ],
      providers: [
        KYCDocumentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateKYCDocumentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});