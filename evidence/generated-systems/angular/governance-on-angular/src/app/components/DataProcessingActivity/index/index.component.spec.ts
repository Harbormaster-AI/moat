
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDataProcessingActivityComponent } from './index.component';
import { DataProcessingActivityService } from '../../../services/DataProcessingActivity.service';

describe('IndexDataProcessingActivityComponent', () => {
  let component: IndexDataProcessingActivityComponent;
  let fixture: ComponentFixture<IndexDataProcessingActivityComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDataProcessingActivityComponent
      ],
      providers: [
        DataProcessingActivityService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDataProcessingActivityComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});