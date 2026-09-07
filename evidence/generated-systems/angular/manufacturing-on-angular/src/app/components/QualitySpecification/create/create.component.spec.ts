
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateQualitySpecificationComponent } from './create.component';
import { QualitySpecificationService } from '../../../services/QualitySpecification.service';
import { Router } from '@angular/router';

describe('CreateQualitySpecificationComponent', () => {
  let component: CreateQualitySpecificationComponent;
  let fixture: ComponentFixture<CreateQualitySpecificationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateQualitySpecificationComponent
      ],
      providers: [
        QualitySpecificationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateQualitySpecificationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});