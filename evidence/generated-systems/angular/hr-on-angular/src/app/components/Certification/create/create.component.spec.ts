
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCertificationComponent } from './create.component';
import { CertificationService } from '../../../services/Certification.service';
import { Router } from '@angular/router';

describe('CreateCertificationComponent', () => {
  let component: CreateCertificationComponent;
  let fixture: ComponentFixture<CreateCertificationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCertificationComponent
      ],
      providers: [
        CertificationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCertificationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});