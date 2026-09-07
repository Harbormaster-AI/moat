
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexMeasureComponent } from './index.component';
import { MeasureService } from '../../../services/Measure.service';

describe('IndexMeasureComponent', () => {
  let component: IndexMeasureComponent;
  let fixture: ComponentFixture<IndexMeasureComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexMeasureComponent
      ],
      providers: [
        MeasureService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexMeasureComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});