
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDataPipelineComponent } from './create.component';
import { DataPipelineService } from '../../../services/DataPipeline.service';
import { Router } from '@angular/router';

describe('CreateDataPipelineComponent', () => {
  let component: CreateDataPipelineComponent;
  let fixture: ComponentFixture<CreateDataPipelineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDataPipelineComponent
      ],
      providers: [
        DataPipelineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDataPipelineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});