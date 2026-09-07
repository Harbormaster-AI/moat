
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditExperimentVariantComponent } from './edit.component';
import { ExperimentVariantService } from '../../../services/ExperimentVariant.service';

describe('EditExperimentVariantComponent', () => {
  let component: EditExperimentVariantComponent;
  let fixture: ComponentFixture<EditExperimentVariantComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditExperimentVariantComponent
      ],
      providers: [
        ExperimentVariantService,
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

    fixture = TestBed.createComponent(EditExperimentVariantComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});