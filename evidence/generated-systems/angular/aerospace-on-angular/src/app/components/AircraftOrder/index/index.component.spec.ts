
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAircraftOrderComponent } from './index.component';
import { AircraftOrderService } from '../../../services/AircraftOrder.service';

describe('IndexAircraftOrderComponent', () => {
  let component: IndexAircraftOrderComponent;
  let fixture: ComponentFixture<IndexAircraftOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAircraftOrderComponent
      ],
      providers: [
        AircraftOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAircraftOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});