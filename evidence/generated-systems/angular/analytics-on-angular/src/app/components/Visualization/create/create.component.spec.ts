
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateVisualizationComponent } from './create.component';
import { VisualizationService } from '../../../services/Visualization.service';
import { Router } from '@angular/router';

describe('CreateVisualizationComponent', () => {
  let component: CreateVisualizationComponent;
  let fixture: ComponentFixture<CreateVisualizationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateVisualizationComponent
      ],
      providers: [
        VisualizationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateVisualizationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});