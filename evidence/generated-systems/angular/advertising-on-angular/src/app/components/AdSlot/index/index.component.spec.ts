
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAdSlotComponent } from './index.component';
import { AdSlotService } from '../../../services/AdSlot.service';

describe('IndexAdSlotComponent', () => {
  let component: IndexAdSlotComponent;
  let fixture: ComponentFixture<IndexAdSlotComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAdSlotComponent
      ],
      providers: [
        AdSlotService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAdSlotComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});