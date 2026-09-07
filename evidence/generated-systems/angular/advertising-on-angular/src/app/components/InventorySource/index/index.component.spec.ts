
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInventorySourceComponent } from './index.component';
import { InventorySourceService } from '../../../services/InventorySource.service';

describe('IndexInventorySourceComponent', () => {
  let component: IndexInventorySourceComponent;
  let fixture: ComponentFixture<IndexInventorySourceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInventorySourceComponent
      ],
      providers: [
        InventorySourceService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInventorySourceComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});