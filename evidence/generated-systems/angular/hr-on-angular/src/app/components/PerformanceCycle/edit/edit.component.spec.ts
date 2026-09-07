
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditPerformanceCycleComponent } from './edit.component';
import { PerformanceCycleService } from '../../../services/PerformanceCycle.service';

describe('EditPerformanceCycleComponent', () => {
  let component: EditPerformanceCycleComponent;
  let fixture: ComponentFixture<EditPerformanceCycleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditPerformanceCycleComponent
      ],
      providers: [
        PerformanceCycleService,
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

    fixture = TestBed.createComponent(EditPerformanceCycleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});