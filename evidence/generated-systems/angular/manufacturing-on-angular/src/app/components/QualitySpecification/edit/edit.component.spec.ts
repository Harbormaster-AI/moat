
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditQualitySpecificationComponent } from './edit.component';
import { QualitySpecificationService } from '../../../services/QualitySpecification.service';

describe('EditQualitySpecificationComponent', () => {
  let component: EditQualitySpecificationComponent;
  let fixture: ComponentFixture<EditQualitySpecificationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditQualitySpecificationComponent
      ],
      providers: [
        QualitySpecificationService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditQualitySpecificationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});