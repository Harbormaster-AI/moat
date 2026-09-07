
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexConversionEventComponent } from './index.component';
import { ConversionEventService } from '../../../services/ConversionEvent.service';

describe('IndexConversionEventComponent', () => {
  let component: IndexConversionEventComponent;
  let fixture: ComponentFixture<IndexConversionEventComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexConversionEventComponent
      ],
      providers: [
        ConversionEventService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexConversionEventComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});