
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDataPipelineComponent } from './index.component';
import { DataPipelineService } from '../../../services/DataPipeline.service';

describe('IndexDataPipelineComponent', () => {
  let component: IndexDataPipelineComponent;
  let fixture: ComponentFixture<IndexDataPipelineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDataPipelineComponent
      ],
      providers: [
        DataPipelineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDataPipelineComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});